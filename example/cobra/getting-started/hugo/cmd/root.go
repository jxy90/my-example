package cmd

import (
	"fmt"
	"github.com/mitchellh/go-homedir"
	"github.com/spf13/cobra"
	"github.com/spf13/cobra/doc"
	"github.com/spf13/viper"
	"os"
	"time"
)

var (
	cfgFile          string
	Verbose          bool
	Source           string
	Region           string
	author           string
	MarkdownDocs     bool
	ReStructuredDocs bool
	ManPageDocs      bool
)

var rootCmd = &cobra.Command{
	Use:   "hugo",
	Short: "Hugo is a very fast static site generator",
	Long: `A Fast and Flexible Static Site Generator built with
                love by spf13 and friends in Go.
                Complete documentation is available at https://gohugo.io`,
	PersistentPreRun: func(cmd *cobra.Command, args []string) {
		fmt.Println("hugo PersistentPreRun")
	},
	PreRun: func(cmd *cobra.Command, args []string) {
		fmt.Println("hugo PreRun")
	},
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("run hugo...")
		fmt.Printf("Verbose: %v\n", Verbose)
		fmt.Printf("Source: %v\n", Source)
		fmt.Printf("Region: %v\n", Region)
		fmt.Printf("Author: %v\n", viper.Get("author"))
		fmt.Printf("Config: %v\n", viper.AllSettings())
	},
	PostRun: func(cmd *cobra.Command, args []string) {
		fmt.Println("hugo PostRun")
	},
	PersistentPostRun: func(cmd *cobra.Command, args []string) {
		fmt.Println("hugo PersistentPostRun")
	},
	TraverseChildren: true,
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

func init() {
	//使用 cobra.OnInitialize() 来初始化配置文件
	//cobra.OnInitialize(initConfig)
	rootCmd.Flags().StringVarP(&cfgFile, "config", "c", "", "config file")
	// 持久标志
	rootCmd.PersistentFlags().BoolVarP(&Verbose, "verbose", "v", false, "verbose output")
	// 本地标志
	rootCmd.Flags().StringVarP(&Source, "source", "s", "", "Source directory to read from")
	// 必选标志
	rootCmd.Flags().StringVarP(&Region, "region", "r", "", "AWS region (required)")
	//rootCmd.MarkFlagRequired("region")

	// 绑定标志到 Viper
	rootCmd.PersistentFlags().StringVar(&author, "author", "YOUR NAME", "Author name for copyright attribution")
	if err := viper.BindPFlag("author", rootCmd.PersistentFlags().Lookup("author")); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	// 生成文档
	rootCmd.Flags().BoolVarP(&MarkdownDocs, "md-docs", "m", false, "gen Markdown docs")
	rootCmd.Flags().BoolVarP(&ReStructuredDocs, "rest-docs", "t", false, "gen ReStructured Text docs")
	rootCmd.Flags().BoolVarP(&ManPageDocs, "man-docs", "a", false, "gen Man Page docs")

	// 定制使用 `help` 命令查看帮助信息输出结果
	//rootCmd.SetHelpCommand(&cobra.Command{
	//	Use:    "help",
	//	Short:  "Custom help command",
	//	Hidden: true,
	//	Run: func(cmd *cobra.Command, args []string) {
	//		fmt.Println("Custom help command")
	//	},
	//})

	// 定制使用 `-h/--help` 命令查看帮助信息输出结果
	//rootCmd.SetHelpFunc(func(command *cobra.Command, strings []string) {
	//	fmt.Println(strings)
	//})

}
func initConfig() {
	if cfgFile != "" {
		viper.SetConfigFile(cfgFile)
	} else {
		home, err := homedir.Dir()
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}

		viper.AddConfigPath(home)
		viper.SetConfigName(".cobra")
	}

	if err := viper.ReadInConfig(); err != nil {
		fmt.Println("Can't read config:", err)
		os.Exit(1)
	}
}

func GenDocs() {
	var err error
	switch {
	case MarkdownDocs:
		err = doc.GenMarkdownTree(rootCmd, "./docs/md")
	case ReStructuredDocs:
		err = doc.GenReSTTree(rootCmd, "./docs/rest")
	case ManPageDocs:
		t := time.Now()
		header := &doc.GenManHeader{
			Title:   "hugo",
			Section: "1",
			Manual:  "hugo Manual",
			Source:  "hugo source",
			Date:    &t,
		}
		err = doc.GenManTree(rootCmd, header, "./docs/man")
	}
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
