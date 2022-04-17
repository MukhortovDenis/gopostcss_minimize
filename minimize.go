package main

import (
	"errors"
	"strings"

	parse "github.com/MukhortovDenis/gopostcss_parser"
)

func Run(ast *parse.AST) error {
	if ast == nil {
		return errors.New("no point for ast")
	}
	for i := range ast.Tokens {
		for k := range ast.Tokens[i].Rules {
			for ik := range *ast.Tokens[i].Rules[k] {
				rule := *ast.Tokens[i].Rules[k]
				*rule[ik] = strings.TrimSuffix(*rule[ik], ",")
			}
		}
	}
	return nil
}
