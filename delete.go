package main

import (
	"fmt"
	"github.com/urfave/cli/v2"
	"io/ioutil"
	"net/url"
	"regexp"
	"strings"
)

func Delete(c *cli.Context) error {
	client, err := GetHttpClient()
	if err != nil {
		return err
	}
	repositoriesList := strings.Split(Account.Repositories, ",")
	for _, repositories := range repositoriesList {
		bundle, err := getBundle(client, repositories)
		if err != nil {
			return err
		}
		err = deleteRepositories(client, bundle, repositories)
		if err != nil {
			return err
		}
	}
	return nil
}

func deleteRepositories(client *httpClient, bundle, repositories string) error {
	reqUrl := fmt.Sprintf("%s/%s/%s", Domain, Account.UserName, repositories)
	param := make(url.Values)
	param.Set("_method", "DELETE")
	param.Set("_pavise[bundle]", bundle)
	param.Set("_pavise[password]", Account.Password)
	response, err := client.Delete(reqUrl, HttpTypePost3W, param)
	if err != nil {
		return err
	}
	data, err := ioutil.ReadAll(response.Body)
	if err != nil {
		return err
	}
	bundleReg := regexp.MustCompile(`success`)
	result := bundleReg.FindSubmatch(data)
	ret := "fail"
	if len(result) > 0 {
		ret = "success"
	}
	fmt.Printf("repositories[%s/%s] delete %s! \n", Account.UserName, repositories, ret)
	return nil
}

func getBundle(client *httpClient, repositories string) (string, error) {
	reqUrl := fmt.Sprintf("%s/%s/%s", Domain, Account.UserName, repositories)
	param := make(url.Values)
	param.Set("path_with_namespace", Account.UserName+"/"+repositories)
	response, err := client.Delete(reqUrl, HttpTypePost3W, param)
	if err != nil {
		return "", err
	}
	data, err := ioutil.ReadAll(response.Body)
	if err != nil {
		return "", err
	}
	bundleReg := regexp.MustCompile(`\[bundle\]\\" type=\\"hidden\\" value=\\"([a-zA-Z0-9=-]*)`)
	result := bundleReg.FindSubmatch(data)
	bundle := ""
	if len(result) > 1 {
		bundle = string(result[1])
	}
	return bundle, nil
}
