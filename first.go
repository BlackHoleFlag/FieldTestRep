package main

import (
	"math"
    "fmt"
    "cryptorandomizer"
)
var(
	raceAll = [6] string {"ork", "elf", "dwarf", "goblin", "human", "undead"}
	nicknameAll = [5] string {"Maki", "Julex", "Alonso", "Trish", "Donkin"}
	orkSkillsAll = [3] skillsStruct{{"mighty fist",5,0} , {"headbash",15,0} , {"body slam", 10,0}}
	elfSkillsAll = [3] skillsStruct{{"sharp vengence", 5,0} , {"fast arrow", 12,0} , {"selfheal", 0, 12}}
	dwarfSkillsAll = [3] skillsStruct{{"hammertime", 7,0} , {"shield bash", 10,0} , {"skill \"Try hard\"", 12,0}}
	goblinSkillsAll = [3] skillsStruct{{"boot stabbing", 4,0} , {"SUCCESSFULL BACK STABBING", 20,0} , {"licking wounds behind a rock", 0,6}}
	humanSkillsAll = [3] skillsStruct{{"falcon PUNCH", 6,0} , {"crane KICK", 8,0 } , {"heal potion", 0,15 }}
	undeadSkillsAll = [3] skillsStruct{{"soul sucking", 5,5}, {"banshi scream", 8,0}, {"powerfull sucking", 10,6}}
)
func main(){
	
	challenger1:= challengerCreator()
	fmt.Println(challenger1)
	challenger1 = challengerModifier(challenger1)
	fmt.Println(challenger1)
	challenger2:= challengerModifier(challengerCreator())
	var winBool bool
	var rounds int
	winStrike := 0
	rounds, winBool = fight(challenger1, challenger2)
	fmt.Println(rounds)
	if winBool == true{
		winStrike = 1
		fmt.Printf("Mighty %v, you pass qualifiers. Now go to Finals, where you going to fight four more warriors\n",challenger1.race)
		for winStrike <5{
			challenger2 = challengerModifier(challengerCreator())
			rounds, winBool = fight(challenger1, challenger2)
			winStrike++
			switch winBool{
			case false: fmt.Printf("Your win streak is %v, but you loose to a stronger opponent\n", winStrike)
			default:
			}
			switch winStrike{
			case 5: fmt.Println("You did it! You are champion of today!")
			default:
			}	
		}
	}
}

type skillsStruct struct{
	name string
	skillDmg, skillHeal float64
}
type challengerStruct struct{
	nickname, race string
	hp, dmg float64
	skill [3]skillsStruct 
	
}

func challengerCreator()challengerStruct{
	a:= cryptorandomizer.Num(5)
	b:= cryptorandomizer.Num(6)
	c:= 25 + float64(cryptorandomizer.Num(50))
	d:= 5 + float64(cryptorandomizer.Num(10))
	var e [3]skillsStruct
	switch b{
	case 0:  e=orkSkillsAll
	case 1:  e=elfSkillsAll
	case 2:  e=dwarfSkillsAll
	case 3:  e=goblinSkillsAll
	case 4:  e=humanSkillsAll
	default: e=undeadSkillsAll
	}
	challenger:= challengerStruct{nicknameAll[a], raceAll[b], c , d, e }
	return challenger
}

func challengerModifier(c challengerStruct)challengerStruct{
	switch c.race{
	case "ork":
		c.hp  = c.hp*1.6
		c.dmg = c.dmg*1.3
	case "goblin":
		c.hp  = c.hp*1.4
		c.dmg = c.dmg*1.35
	case "dwarf":
		c.hp  = c.hp*1.7
		c.dmg = c.dmg*1.25
	case "elf":
		c.hp  = c.hp*1.35
		c.dmg = c.dmg*1.5
	case "human":
		c.hp  = c.hp*1.45
		c.dmg = c.dmg*1.45
	default:
		c.hp  = c.hp*1.8
		c.dmg = c.dmg*0.9
	}
	c.hp = roundUp(c.hp)
	c.dmg = roundUp(c.dmg)
	return c
}
func fight(opponent1, opponent2 challengerStruct)(int, bool){
	var winner bool
	rounds:=0
    var x challengerStruct = opponent1
    var y challengerStruct = opponent2
    for x.hp > 0 && y.hp > 0 {
        rounds++
        //who act first
       whoActFirst:= cryptorandomizer.Num(2)
        if whoActFirst == 0 { 
            whatSkill1 := cryptorandomizer.Num(3)
            whatSkill2 := cryptorandomizer.Num(3)
            fmt.Printf("%v the %v using %v. ", x.nickname, x.race, x.skill[whatSkill1].name)
			y.hp = y.hp - x.skill[whatSkill1].skillDmg
			x.hp = x.hp + x.skill[whatSkill1].skillHeal
            fmt.Printf("%v the %v have %v health left\n", y.nickname, y.race, y.hp )
            if y.hp <0 {
                fmt.Println("Now iam gonna use my FINAL MOVE")
            }
            fmt.Printf("%v the %v using %v. ", y.nickname, y.race, y.skill[whatSkill2].name)
			x.hp = x.hp - y.skill[whatSkill2].skillDmg
			y.hp = y.hp + y.skill[whatSkill2].skillHeal
            fmt.Printf("%v the %v have %v health left\n", x.nickname, x.race, x.hp)
        }else{
            whatSkill1 := cryptorandomizer.Num(2)
            whatSkill2 := cryptorandomizer.Num(2)
            fmt.Printf("%v the %v using %v. ", y.nickname, y.race, y.skill[whatSkill2].name)
			x.hp = x.hp - y.skill[whatSkill1].skillDmg
			y.hp = y.hp + y.skill[whatSkill1].skillHeal
            fmt.Printf("%v the %v have %v health left\n", x.nickname, x.race, x.hp)
            if x.hp < 0 {
                fmt.Println("Now iam gonna use my FINAL MOVE")}
            fmt.Printf("%v the %v using %v. ", x.nickname, x.race, x.skill[whatSkill1].name)
			y.hp = y.hp - x.skill[whatSkill2].skillDmg
			x.hp = x.hp + x.skill[whatSkill2].skillHeal
            fmt.Printf("%v the %v have %v health left\n", y.nickname, y.race, y.hp)
        
        }
	 }
	 if x.hp>0 {
		winner = true
		fmt.Printf("You win!, %v %v gracefully lost.\n", y.nickname , y.race)
	 }else{
		 winner = false
		 fmt.Printf("You lose, %v was too strong\n", y.race )
	 }
     return rounds, winner
}

func roundUp(input float64)float64{
	var round float64
	pow:=math.Pow(10, float64(2))
	digit:= pow * input
	round = math.Ceil(digit)
	newVal:= round/pow
	return newVal
}