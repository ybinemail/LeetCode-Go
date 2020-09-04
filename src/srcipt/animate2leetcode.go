package main

import (
	"fmt"
	"io"
	"io/ioutil"
	"os"
	"path/filepath"
	"regexp"
	"strings"
)

const ANIMATION = "D:\\WorkspaceM\\GoProject\\LeetCodeAnimation"

//const ANIMATION = "D:\\WorkspaceM\\GoProject\\LeetCodeAnimation-1"

const LEETCODE = "D:\\WorkspaceM\\GoProject\\LeetCode-Go\\leetcode"

//const LEETCODE = "D:\\test"

const REGEXPSTR = "^[\\d]{4}\\D"

func main() {
	animationFiles, _ := ioutil.ReadDir(ANIMATION)

	leetcodeFiles, _ := ioutil.ReadDir(LEETCODE)

	reg, err := regexp.Compile(REGEXPSTR)
	if err != nil {
		fmt.Println("regexp is error")
	}

	// leetMap
	leetMap := make(map[string]string)
	for _, leetcode := range leetcodeFiles {
		if reg.MatchString(leetcode.Name()) {
			codeIndex := leetcode.Name()[0:4]
			leetMap[codeIndex] = leetcode.Name()
		}
	}
	fmt.Printf("get leetcode dir count: %d\n", len(leetMap))
	fmt.Println("===================== start get animate file =========================")
	for _, animate := range animationFiles {
		if reg.MatchString(animate.Name()) {
			animateCode := animate.Name()[0:4]
			leetCodeFile, ok := leetMap[animateCode]
			if !ok {
				fmt.Printf("===> animate %s : %s is not in leetcode \n", animateCode, animate.Name())
				// 不存在，创建文件对应目录
				leetCodeFile = strings.Replace(animate.Name(), "-", ".", 1)
				newLeecode := LEETCODE + string(os.PathSeparator) + leetCodeFile

				os.Mkdir(newLeecode, os.ModeDir)
			} else {
				fmt.Printf("<==== animate %s : %s exist in leetcode \n", animateCode, animate.Name())
			}
			fmt.Printf("---> start copy 【%s】 to 【%s】\n", animate.Name(), leetCodeFile)
			fmt.Println()
			dirPath := LEETCODE + string(os.PathSeparator) + leetCodeFile
			srcPath := ANIMATION + string(os.PathSeparator) + animate.Name()
			copyDir(srcPath, dirPath)
		}
	}
}
func copyDir(src string, dest string) {
	src_original := src
	err := filepath.Walk(src, func(src string, f os.FileInfo, err error) error {
		if f == nil {
			return err
		}
		if f.IsDir() {
			fmt.Println(f.Name())
			copyDir(f.Name(), dest+"/"+f.Name())
		} else {
			fmt.Println(src)
			fmt.Println(src_original)
			fmt.Println(dest)
			dest_new := strings.Replace(src, src_original, dest, -1)
			fmt.Println(dest_new)
			fmt.Println("CopyFile:" + src + " to " + dest_new)
			CopyFile(src, dest_new)
		}
		//println(path)
		return nil
	})
	if err != nil {
		fmt.Printf("filepath.Walk() returned %v\n", err)
	}
}

//egodic directories
func getFilelist(path string) {
	err := filepath.Walk(path, func(path string, f os.FileInfo, err error) error {
		if f == nil {
			return err
		}
		if f.IsDir() {
			return nil
		}
		println(path)
		return nil
	})
	if err != nil {
		fmt.Printf("filepath.Walk() returned %v\n", err)
	}
}
func PathExists(path string) (bool, error) {
	_, err := os.Stat(path)
	if err == nil {
		return true, nil
	}
	if os.IsNotExist(err) {
		return false, nil
	}
	return false, err
}

//copy file
func CopyFile(src, dst string) (w int64, err error) {
	srcFile, err := os.Open(src)
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	defer srcFile.Close()
	fmt.Println("dst:" + dst)
	dst_slices := strings.Split(dst, "\\")
	dst_slices_len := len(dst_slices)
	dest_dir := ""
	for i := 0; i < dst_slices_len-1; i++ {
		dest_dir = dest_dir + dst_slices[i] + "\\"
	}
	//dest_dir := getParentDirectory(dst)
	fmt.Println("dest_dir:" + dest_dir)
	b, err := PathExists(dest_dir)
	if b == false {
		err := os.Mkdir(dest_dir, os.ModePerm) //在当前目录下生成md目录
		if err != nil {
			fmt.Println(err)
		}
	}
	dstFile, err := os.Create(dst)

	if err != nil {
		fmt.Println(err.Error())
		return
	}

	defer dstFile.Close()

	return io.Copy(dstFile, srcFile)
}
