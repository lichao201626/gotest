package main

import (
	"fmt"
	"log"
	"strconv"
)

func split(str string) []string {
	log.Println("before split", str)
	res := []string{}
	fmt.Println("before split", str)
	temp := ""
	for k, v := range str {
		x := string(v)
		log.Println(k, v, x)

		switch x {
		case " ":
			break
		case Plus:
			res = append(res, temp)
			res = append(res, Plus)
			temp = ""
			break
		case Minus:
			res = append(res, temp)
			res = append(res, Minus)
			temp = ""
			break
		case Multiple:
			res = append(res, temp)
			res = append(res, Multiple)
			temp = ""
			break
		case Divide:
			res = append(res, temp)
			res = append(res, Divide)
			temp = ""
			break
		default:
			temp += x
		}
		log.Println("temp", temp)
	}

	//res := strings.Split(str, Plus)
	res = append(res, temp)
	return res
}

func convert(arr []string) {

}

func eval(s string) int {
	sa := split(s)
	vstack := inits()
	ostack := inits()
	for k := len(sa) - 1; k >= 0; k-- {
		log.Println("kkkkk", k)
		switch sa[k] {
		case Plus:
			ostack.push(Plus)
			break
		case Minus:
			ostack.push(Minus)
			break
		case Multiple:
			v1 := vstack.pop()
			s := ""
			switch v1.Value.(type) {
			case string:
				s = v1.Value.(string)
			}
			k--
			v2 := sa[k]
			av := cal(v2, s, sa[k+1])
			vstack.push(strconv.Itoa(av))
			break
		case Divide:
			v1 := vstack.pop()
			s := ""
			switch v1.Value.(type) {
			case string:
				s = v1.Value.(string)
			}
			k--
			v2 := sa[k]
			av := cal(v2, s, sa[k+1])
			vstack.push(strconv.Itoa(av))
			break
		default:
			vstack.push(sa[k])
			break
		}
	}
	//vstack.show()
	//ostack.show()
	x1 := vstack.pop()
	tmp1 := ""
	switch x1.Value.(type) {
	case string:
		tmp1 = x1.Value.(string)
	}
	res := 0
	log.Println("tmp1", tmp1, vstack.Size)
	for vstack.Size > 0 && ostack.Size > 0 {
		x2 := vstack.pop()
		log.Println("tmp1", x2, x2.Value)
		tmp2 := ""
		switch x2.Value.(type) {
		case string:
			tmp2 = x2.Value.(string)
		}
		log.Println("tmp2", tmp2)
		o := ostack.pop()
		tmp3 := ""
		switch o.Value.(type) {
		case string:
			tmp3 = o.Value.(string)
		}
		res = cal(tmp1, tmp2, tmp3)
		tmp1 = strconv.Itoa(res)
	}
	return res
}

func cal(s1 string, s2 string, oper string) int {
	log.Println("before cal", s1, s2, oper)
	res := 0
	v1, err1 := strconv.Atoi(s1)
	v2, err2 := strconv.Atoi(s2)
	if err1 == nil && err2 == nil {
		switch oper {
		case Plus:
			res = v1 + v2
			break
		case Minus:
			res = v1 - v2
			break
		case Multiple:
			res = v1 * v2
			break
		case Divide:
			res = v1 / v2
			break
		default:
			break
		}
	}
	log.Println("after cal", v1, v2, oper, res)
	return res
}
