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
)

type ArgType struct {
	Flag string
	Type string
}

func listObjs(cmd *cobra.Command, args []string) error {
	env.InitByConfig(flag_config)
	rootUrl := env.Config().GetMapString("apiServer", "url", "http://localhost:3330/api")
	class := cmd.Parent().Aliases[0]
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
	fmt.Println(resBody)
	return nil
}

func getObjFunc(assoc []string) func(cmd *cobra.Command, args []string) error {
	return func(cmd *cobra.Command, args []string) error {
		return getObj(cmd, assoc)
	}
}

func getObj(cmd *cobra.Command, assoc []string) error {
	if flag_id == 0 {
		return errors.New("--id -i is null")
	}
	env.InitByConfig(flag_config)

	rootUrl := env.Config().GetMapString("apiServer", "url", "http://localhost:3330/api")
	class := cmd.Parent().Aliases[0]
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

func createObj(cmd *cobra.Command, args []string) error {
	if flag_data_path == "" {
		return errors.New("--Data -D is null")
	}

	env.InitByConfig(flag_config)

	rootUrl := env.Config().GetMapString("apiServer", "url", "http://localhost:3330/api")
	class := cmd.Parent().Aliases[0]
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

func removeObj(cmd *cobra.Command, args []string) error {
	if flag_id == 0 {
		return errors.New("--id -i is null")
	}
	env.InitByConfig(flag_config)

	rootUrl := env.Config().GetMapString("apiServer", "url", "http://localhost:3330/api")
	class := cmd.Parent().Aliases[0]
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

func invokeObjFunc(method string, argTypes []ArgType) func(cmd *cobra.Command, args []string) error {
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
		return invokeObj(cmd, method, params)
	}
}

func invokeObj(cmd *cobra.Command, method string, args []interface{}) error {
	if flag_id == 0 {
		return errors.New("--id -i is null")
	}
	env.InitByConfig(flag_config)

	rootUrl := env.Config().GetMapString("apiServer", "url", "http://localhost:3330/api")
	class := cmd.Parent().Aliases[0]

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
	resBody := res.Body()
	if res.StatusCode() != 200 {
		return fmt.Errorf("%s is failed. %s", method, resBody)
	}
	fmt.Println(string(resBody))
	return nil
}





