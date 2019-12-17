package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
	"os"
	"qovery.go/api"
	"qovery.go/util"
)

var checkoutCmd = &cobra.Command{
	Use:   "checkout",
	Short: "Equivalent to 'git checkout' but with Qovery magic sauce",
	Long: `CHECKOUT performs 'git checkout' action and set Qovery properties to target the right environment . For example:

	qovery checkout`,
	Run: func(cmd *cobra.Command, args []string) {
		if len(args) != 1 {
			fmt.Println("qovery checkout <branch>")
			os.Exit(1)
		}

		api.DeleteLocalConfiguration()

		branch := args[0]
		branchName := util.CurrentBranchName()
		projectName := util.CurrentQoveryYML().Application.Project

		if branchName == "" || projectName == "" {
			fmt.Println("The current directory is not a Qovery project. Please consider using 'qovery init'")
			os.Exit(1)
		}

		project := api.GetProjectByName(projectName)
		if project == nil {
			fmt.Println("The project does not exist. Are you well authenticated with the right user? Do 'qovery auth' to be sure")
			os.Exit(1)
		}

		// checkout branch
		util.Checkout(branch)

		applications := api.ListApplicationsRaw(project.Id, branchName)
		api.SaveLocalConfiguration(applications)
	},
}

func init() {
	RootCmd.AddCommand(checkoutCmd)
}
