package main

import (
	"bytes"
	"flag"
	"fmt"
	"io"
	"io/ioutil"
	"os"
	"os/exec"
	"path/filepath"

	"github.com/pkg/errors"
)

func isParamsVlaid(output string, name string, description string) (bool, error) {
	if output == "" {
		fmt.Println("output 為必要參數，不可為空")
		return false, nil
	}
	if name == "" {
		fmt.Println("name 為必要參數，不可為空")
		return false, nil
	}
	fileInfo, err := os.Stat(output)
	if err != nil {
		if os.IsNotExist(err) {
			fmt.Printf("輸出路徑 '%s' 不存在\n", output)
			return false, nil
		} else {
			return false, errors.Wrap(err, "讀取輸出路徑資訊時發生錯誤")
		}
	}
	if !fileInfo.IsDir() {
		fmt.Printf("輸出路徑 '%s' 不是資料夾\n", output)
		return false, nil
	}
	return true, nil
}

func copyProject(folder string, name string, description string) error {
	_, err := os.Stat(folder)
	if err != nil {
		if !os.IsNotExist(err) {
			return errors.Wrap(err, "讀取輸出路徑資訊時發生錯誤")
		}
		// 資料夾不存在，協助生成
		err = os.Mkdir(folder, os.ModePerm)
		if err != nil {
			return errors.Wrap(err, "生成資料夾時發生錯誤")
		}
	}
	// public
	copyFolder("./seed/public", filepath.Join(folder, "public"))
	// view
	copyFolder("./seed/views", filepath.Join(folder, "views"))
	path := filepath.Join(folder, "views/layouts/main.hbs")
	err = modifyFile(path,
		"<title>Project Title</title>",
		fmt.Sprintf("<title>%s</title>", name))
	if err != nil {
		fmt.Printf("修改專案名稱時發生錯誤(%s)\n", path)
	}
	// index.js
	copyFile("./seed/index.js", filepath.Join(folder, "index.js"))
	// package.json
	path = filepath.Join(folder, "package.json")
	copyFile("./seed/package.json", path)
	err = modifyFile(path,
		`"name": "express-project-seed"`,
		fmt.Sprintf(`"name": "%s"`, name))
	if err != nil {
		fmt.Printf("修改專案名稱時發生錯誤(%s)\n", path)
	}
	if description != "" {
		err = modifyFile(path,
			`"description": "初始化 Express 所需專案與資源，快速開啟一個新的 Express 專案"`,
			fmt.Sprintf(`"description": "%s"`, description))
		if err != nil {
			fmt.Printf("修改專案描述時發生錯誤(%s)\n", path)
		}
	}
	// package-lock.json
	copyFile("./seed/package-lock.json", filepath.Join(folder, "package-lock.json"))
	// .gitignore
	copyFile("./seed/.gitignore", filepath.Join(folder, ".gitignore"))
	return nil
}

func copyFolder(src, dst string) error {
	err := filepath.Walk(src, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}

		relativePath, err := filepath.Rel(src, path)
		if err != nil {
			return err
		}

		destPath := filepath.Join(dst, relativePath)

		if info.IsDir() {
			err := os.MkdirAll(destPath, os.ModePerm)
			if err != nil {
				return err
			}
		} else {
			err := copyFile(path, destPath)
			if err != nil {
				return err
			}
		}

		return nil
	})

	return err
}

func copyFile(src, dst string) error {
	srcFile, err := os.Open(src)
	if err != nil {
		return err
	}
	defer srcFile.Close()

	dstFile, err := os.Create(dst)
	if err != nil {
		return err
	}
	defer dstFile.Close()
	_, err = io.Copy(dstFile, srcFile)
	if err != nil {
		return err
	}
	return nil
}

func modifyFile(filePath, oldContent, newContent string) error {
	content, err := ioutil.ReadFile(filePath)
	if err != nil {
		return errors.Wrap(err, "讀取檔案時發生錯誤")
	}
	// 修改文件内容
	modifiedContent := bytes.Replace(content, []byte(oldContent), []byte(newContent), -1)
	// 將內容寫回檔案
	err = ioutil.WriteFile(filePath, modifiedContent, os.ModePerm)
	if err != nil {
		return errors.Wrap(err, "寫入檔案時發生錯誤")
	}
	return nil
}

func main() {
	var output string
	var name string
	var description string

	flag.StringVar(&output, "output", "", "專案輸出的資料夾")
	flag.StringVar(&name, "name", "", "專案名稱")
	flag.StringVar(&description, "description", "", "Enable verbose mode")

	// 解析命令列參數
	flag.Parse()

	// 檢查參數是否有效，若無效則直接結束
	isValid, err := isParamsVlaid(output, name, description)
	if err != nil {
		fmt.Printf("檢查參數時發生錯誤, err:\n%+v\n", err)
		return
	}
	if !isValid {
		fmt.Println("傳入無效參數")
		return
	}
	folder := filepath.Join(output, name)
	err = copyProject(folder, name, description)
	if err != nil {
		fmt.Printf("複製專案時發生錯誤, err:\n%+v\n", err)
		return
	}
	cmd := exec.Command("npm", "install")
	cmd.Dir = folder
	err = cmd.Run()
	if err != nil {
		fmt.Printf("安裝依賴套件時發生錯誤, err:\n%+v\n", err)
		return
	}
	fmt.Printf("完成 Express 專案 %s 的初始化", name)
}
