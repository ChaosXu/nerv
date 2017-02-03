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
	body := res.Body()
	fmt.Println(string(body))
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
	body := res.Body()
	fmt.Println(string(body))
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
	resBody := res.Body()

	fmt.Println(string(resBody))
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
	body := res.Body()
	fmt.Println(string(body))
	return nil
	return nil
}

func invokeObjFunc(method string, args ...string) func(cmd *cobra.Command, args []string) error {
	return func(cmd *cobra.Command, args []string) error {
		cmd.Flags().Get
		return invokeObj(cmd, method, args)
	}
}

func invokeObj(cmd *cobra.Command, method string, args ...interface{}) error {
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

	res, err := resty.R().
			SetHeader("Content-Type", "application/json").
			SetBody(body).
			Post(fmt.Sprintf("%s/objs/%s/%d/%s", rootUrl, class, flag_id, method))
	resBody := res.Body()
	fmt.Println(string(resBody))
	return nil
}





