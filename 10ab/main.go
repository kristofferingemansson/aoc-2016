package main

import (
	"os"
	"log"
	"io/ioutil"
	"fmt"
	"strings"
	"regexp"
	"strconv"
)

const (
	TARGET_BOT = "bot"
	TARGET_OUTPUT = "output"
)

type BotId int
type ChipId int
type OutputId int

type Bot struct {
	id BotId
	a ChipId
	b ChipId
}

type BotList map[BotId]Bot

type Rule struct {
	botLow BotId
	botHigh BotId
	outputLow OutputId
	outputHigh OutputId
}

type RuleMap map[BotId]Rule

type OutputMap map[OutputId][]ChipId

var (
	RegexpState, _ = regexp.Compile("value (\\d+) goes to bot (\\d+)")
	RegexpRule, _ = regexp.Compile("bot (\\d+) gives low to (bot|output) (\\d+) and high to (bot|output) (\\d+)")
)

func main() {
	data := GetInputData()
	rules, bots := ReadInitialState(data)
	outputs := OutputMap{}

	for true {
		changes := Tick(rules, bots, outputs)
		if changes == 0 {
			break;
		}
	}

	fmt.Println("Rules: ", rules)
	fmt.Println("Bots: ", bots)
	fmt.Println("Outputs: ", outputs)
}

func GetInputData() []string {
	dir, err := os.Getwd()
	if err != nil {
		log.Fatalln(err)
	}

	input, err := ioutil.ReadFile(dir + "/10ab/input.txt")
	if err != nil {
		log.Fatalln(err)
	}

	return strings.Split(strings.Trim(string(input), "\n"), "\n")
}

func ReadInitialState(rows []string) (RuleMap, BotList) {
	bots := BotList{}
	rules := RuleMap{}

	for _, row := range rows {
		stateMatch := RegexpState.FindStringSubmatch(row)
		if stateMatch != nil {
			botId := StringToBotId(stateMatch[2])
			bot, found := bots[botId]
			if !found {
				bot = Bot{id: botId}
			}
			bot.Take(StringToChipId(stateMatch[1]))
			bots[botId] = bot
			continue
		}

		ruleMatch := RegexpRule.FindStringSubmatch(row)
		if ruleMatch != nil {
			botId := StringToBotId(ruleMatch[1])
			rule, found := rules[botId]
			if !found {
				rule = Rule{botLow: -1, botHigh: -1, outputLow: -1, outputHigh: -1}
			}

			if ruleMatch[2] == TARGET_BOT {
				rule.botLow = StringToBotId(ruleMatch[3])
			} else {
				rule.outputLow = StringToOutputId(ruleMatch[3])
			}

			if ruleMatch[4] == TARGET_BOT {
				rule.botHigh = StringToBotId(ruleMatch[5])
			} else {
				rule.outputHigh = StringToOutputId(ruleMatch[5])
			}

			rules[botId] = rule
		}
	}

	return rules, bots
}

func Tick(rules RuleMap, bots BotList, outputMap OutputMap) int {
	changes := 0
	for botId, bot := range bots {
		if bot.NumChips() == 2 {
			rule, found := rules[botId]
			if !found {
				continue
			}
			changes++
			if rule.outputLow != -1 {
				outputMap.Take(rule.outputLow, bot.GiveLow())
			}
			if rule.outputHigh != -1 {
				outputMap.Take(rule.outputHigh, bot.GiveHigh())
			}
			if rule.botLow != -1 {
				targetBot, found := bots[rule.botLow]
				if !found {
					targetBot = Bot{id: rule.botLow}
				}
				targetBot.Take(bot.GiveLow())
				bots[rule.botLow] = targetBot
			}
			if rule.botHigh != -1 {
				targetBot, found := bots[rule.botHigh]
				if !found {
					targetBot = Bot{id: rule.botHigh}
				}
				targetBot.Take(bot.GiveHigh())
				bots[rule.botHigh] = targetBot
			}
			bots[botId] = bot
		}
	}
	return changes
}

func (b *Bot) Take(c ChipId) {
	if b.a == 0 {
		b.a = c
	} else if b.b == 0 {
		b.b = c
	} else {
		b.a = b.b
		b.b = c
	}

	if b.a == 61 && b.b == 17 || b.a == 17 && b.b == 61 {
		fmt.Printf("Bot responsible for 61+17: %v\n", b)
	}
}

func (b *Bot) GiveLow() ChipId {
	if b.a < b.b || b.b == 0 {
		ret := b.a
		b.a = 0
		return ret
	} else if b.b < b.a || b.a > 0 {
		ret := b.b
		b.b = 0
		return ret
	}
	return 0
}

func (b *Bot) GiveHigh() ChipId {
	if b.a > b.b {
		ret := b.a
		b.a = 0
		return ret
	}
	ret := b.b
	b.b = 0
	return ret
}

func (b *Bot) NumChips() int {
	ret := 0
	if b.a > 0 {
		ret++
	}
	if b.b > 0 {
		ret++
	}
	return ret
}

func (o *OutputMap) Take(outputId OutputId, chipId ChipId) {
	output, found := (*o)[outputId]
	if !found {
		output = []ChipId{}
	}
	output = append(output, chipId)
	(*o)[outputId] = output
}

func StringToChipId(s string) ChipId {
	x, _ := strconv.Atoi(s)
	return ChipId(x)
}

func StringToBotId(s string) BotId {
	x, _ := strconv.Atoi(s)
	return BotId(x)
}

func StringToOutputId(s string) OutputId {
	x, _ := strconv.Atoi(s)
	return OutputId(x)
}