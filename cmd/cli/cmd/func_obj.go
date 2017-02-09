package cmd

import (
	"fmt"
	"errors"
	"github.com/spf13/cobra"
	"github.com/go-resty/resty"
	"github.com/ChaosXu/nerv/lib/env"
	"github.com/toolkits/file"
	"strings"
	"encoding/json"
	"github.com/ChaosXu/nerv/lib/cli/format"
)

type ArgType struct {
	Flag string
	Type string
}

func listObjsFunc(class string, format *format.Page) func(cmd *cobra.Command, args []string) error {
	return func(cmd *cobra.Command, args []string) error {
		env.InitByConfig(flag_config)
		rootUrl := env.Config().GetMapString("apiServer", "url", "http://localhost:3330/api")
		class := class
		url := fmt.Sprintf("%s/objs/%s", rootUrl, class)

		res, err := resty.R().
				SetHeader("Content-Type", "application/json").
				Get(url)
		if err != nil {
			return err
		}
		resBody := string(res.Body())
		if res.StatusCode() != 200 {
			return fmt.Errorf("command is failed. %s", resBody)
		}

		data := map[string]interface{}{}
		if err := json.Unmarshal(res.Body(), &data); err != nil {
			return err
		}

		format.Print(data)
		return nil
	}
}

func getObjFunc(class string, assoc []string) func(cmd *cobra.Command, args []string) error {
	return func(cmd *cobra.Command, args []string) error {
		return getObj(class, cmd, assoc)
	}
}

func getObj(class string, cmd *cobra.Command, assoc []string) error {
	if flag_id == 0 {
		return errors.New("--id -i is null")
	}
	env.InitByConfig(flag_config)

	rootUrl := env.Config().GetMapString("apiServer", "url", "http://localhost:3330/api")

	var url string
	if len(assoc) > 0 {
		ass := strings.Join(assoc, ",")
		url = fmt.Sprintf("%s/objs/%s/%d?associations=%s", rootUrl, class, flag_id, ass)
	} else {
		url = fmt.Sprintf("%s/objs/%s/%d", rootUrl, class, flag_id)
	}

	res, err := resty.R().
			SetHeader("Content-Type", "application/json").
			Get(url)
	if err != nil {
		return err
	}
	resBody := string(res.Body())
	if res.StatusCode() != 200 {
		return fmt.Errorf("command is failed. %s", resBody)
	}
	fmt.Println(resBody)
	return nil
}

func createObjFunc(class string) func(cmd *cobra.Command, args []string) error {
	return func(cmd *cobra.Command, args []string) error {
		if flag_data_path == "" {
			return errors.New("--Data -D is null")
		}

		env.InitByConfig(flag_config)

		rootUrl := env.Config().GetMapString("apiServer", "url", "http://localhost:3330/api")
		url := fmt.Sprintf("%s/objs/%s", rootUrl, class)
		body, err := file.ToTrimString(flag_data_path)
		if err != nil {
			return err
		}
		res, err := resty.R().
				SetHeader("Content-Type", "application/json").
				SetBody(body).
				Post(url)
		if err != nil {
			return err
		}
		resBody := string(res.Body())
		if res.StatusCode() != 200 {
			return fmt.Errorf("command is failed. %s", resBody)
		}
		fmt.Println(resBody)
		return nil
	}
}

func removeObjFunc(class string) func(cmd *cobra.Command, args []string) error {
	return func(cmd *cobra.Command, args []string) error {
		if flag_id == 0 {
			return errors.New("--id -i is null")
		}
		env.InitByConfig(flag_config)

		rootUrl := env.Config().GetMapString("apiServer", "url", "http://localhost:3330/api")
		url := fmt.Sprintf("%s/objs/%s/%d", rootUrl, class, flag_id)

		res, err := resty.R().
				SetHeader("Content-Type", "application/json").
				Delete(url)
		if err != nil {
			return err
		}
		resBody := string(res.Body())
		if res.StatusCode() != 200 {
			return fmt.Errorf("command is failed. %s", resBody)
		}
		fmt.Println(resBody)
		return nil
	}
}

func invokeSvcFunc(class string, method string, argTypes []ArgType) func(cmd *cobra.Command, args []string) error {
	return func(cmd *cobra.Command, args []string) error {
		params := []interface{}{}
		for _, argType := range argTypes {
			switch argType.Type {
			case "string":
				if v, err := cmd.Flags().GetString(argType.Flag); err != nil {
					return err
				} else {
					params = append(params, v)
				}
			case "uint":
				if v, err := cmd.Flags().GetUint(argType.Flag); err != nil {
					return err
				} else {
					params = append(params, v)
				}
			case "ref":
				if v, err := cmd.Flags().GetString(argType.Flag); err != nil {
					return err
				} else {
					buf, err := file.ToBytes(v)
					if err != nil {
						return err
					}
					data := map[string]interface{}{}
					if err := json.Unmarshal(buf, &data); err != nil {
						return err
					} else {
						params = append(params, data)
					}
				}
			default:
				fmt.Errorf("unsupported arg type %s", argType.Type)
			}
		}
		return invokeSvc(class, cmd, method, params)
	}
}

func invokeObjFunc(class string, method string, argTypes []ArgType) func(cmd *cobra.Command, args []string) error {
	return func(cmd *cobra.Command, args []string) error {
		params := []interface{}{}
		for _, argType := range argTypes {
			switch argType.Type {
			case "string":
				if v, err := cmd.Flags().GetString(argType.Flag); err != nil {
					return err
				} else {
					params = append(params, v)
				}
			case "uint":
				if v, err := cmd.Flags().GetUint(argType.Flag); err != nil {
					return err
				} else {
					params = append(params, v)
				}
			default:
				fmt.Errorf("unsupported arg type %s", argType.Type)

			}
		}
		return invokeObj(class, cmd, method, params)
	}
}

func invokeObj(class string, cmd *cobra.Command, method string, args []interface{}) error {
	if flag_id == 0 {
		return errors.New("--id -i is null")
	}
	env.InitByConfig(flag_config)

	rootUrl := env.Config().GetMapString("apiServer", "url", "http://localhost:3330/api")

	b, err := json.Marshal(args)
	if err != nil {
		return err
	}
	body := string(b)
	fmt.Println(body)

	res, err := resty.R().
			SetHeader("Content-Type", "application/json").
			SetBody(body).
			Post(fmt.Sprintf("%s/objs/%s/%d/%s", rootUrl, class, flag_id, method))
	if err != nil {
		return err
	}
	resBody := string(res.Body())
	if res.StatusCode() != 200 {
		return fmt.Errorf("%s is failed. %s", method, resBody)
	}
	fmt.Println(string(resBody))
	return nil
}

func invokeSvc(class string, cmd *cobra.Command, method string, args []interface{}) error {
	env.InitByConfig(flag_config)

	rootUrl := env.Config().GetMapString("apiServer", "url", "http://localhost:3330/api")
	b, err := json.Marshal(args)
	if err != nil {
		return err
	}
	body := string(b)
	fmt.Println(body)
	url := fmt.Sprintf("%s/objs/%s/%s", rootUrl, class, method)
	fmt.Println(url)
	fmt.Println(body)
	res, err := resty.R().
			SetHeader("Content-Type", "application/json").
			SetBody(body).
			Post(url)
	if err != nil {
		return err
	}
	resBody := string(res.Body())
	if res.StatusCode() != 200 {
		return fmt.Errorf("%s is failed. %s", method, resBody)
	}
	fmt.Println(string(resBody))
	return nil
}





