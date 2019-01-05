package cmd

import (
	"fmt"
	"os"
	"log"
	"os/exec"
	"path/filepath"
	"strconv"
)

func executor(terraformTargetDir string, command string) {
	baseDir, err := os.Getwd()
	if err != nil {
		log.Fatal(err)
	}
	
	// Set the target directory terraform will run in.  Defaults to all
	var terraformJob []string
	if terraformTargetDir == "all" {
		terraformJob = executionOrder		
	} else {
		terraformJob = append(terraformJob, terraformTargetDir)
	}

	// Iterate over each terraform directory and run init then plan
	for i := range terraformJob {
		targetDir := filepath.FromSlash(fmt.Sprint(baseDir, "/", terraformJob[i]))
		if err := os.Chdir(targetDir); err != nil {
			log.Fatal(err)
		}

		defer os.Chdir(baseDir)

		os.RemoveAll(".terraform")
		os.RemoveAll("terraform.tfstate.d")

		workspace := fmt.Sprint("TF_WORKSPACE=", terraformJob[i], "-", environment,
			"-uswest2")

		// Terraform init
		initArg1 := fmt.Sprintf("-backend-config=%s", backendConf)

		tfinit := exec.Command("terraform", "init", "-input=false", "-reconfigure", initArg1)
		tfinit.Env = os.Environ()
		tfinit.Env = append(tfinit.Env, workspace)
		tfinitOut, _ := tfinit.CombinedOutput()
		fmt.Printf("combined out:\n%s\n", string(tfinitOut))

		// Terraform plan
		arg1 := fmt.Sprintf("-var-file=%s", strconv.Quote(environmentConf))
		arg2 := "-var"
		arg3 := strconv.Quote(fmt.Sprintf("image=%s", containerImage))

		tfplan := exec.Command("terraform", "plan", "-out=tfplan", "-input=false", arg1, arg2, arg3)
		tfplan.Env = os.Environ()
		tfplan.Env = append(tfplan.Env, workspace)
		tfplanOut, _ := tfplan.CombinedOutput()
		fmt.Printf("combined out:\n%s\n", string(tfplanOut))

		// Terraform apply
		if command == "apply" {
			tfapply := exec.Command("terraform", "apply", "-input=tfplan", "tfplan")
			tfapply.Env = os.Environ()
			tfapply.Env = append(tfapply.Env, workspace)
			tfapplyOut, _ := tfapply.CombinedOutput()
			fmt.Printf("combined out:\n%s\n", string(tfapplyOut))
		}
		
	} 
}
