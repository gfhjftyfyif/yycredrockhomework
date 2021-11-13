package main

import "fmt"

func main() {
	for {
		fmt.Println("请输入技能")
		var x string
		fmt.Scanln(&x)
		switch {
		case x == "龙卷风摧毁停车场":
			ReleaseSkill("龙卷风摧毁停车场", func(skillName1 string) {
				fmt.Println("尝尝我的厉害吧！", skillName1)
			})
		case x == "MK82":
			ReleaseSkill("MK82", func(skillName2 string) {
				fmt.Println("所罗门哟，我回来了！", skillName2)
			})
		case x == "阿福揍扁了成龙":
			ReleaseSkill("阿福揍扁了成龙", func(skillName3 string) {
				fmt.Println("我叫黑虎阿福，你准备受死吧！", skillName3)
			})
		}
	}
}
func ReleaseSkill(skillNames string, releaseSkillFunc func(string)) {
	releaseSkillFunc(skillNames)
}