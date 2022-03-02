package mysql

import (
	"bytes"
	"os"
	"os/exec"

	qmodel "github.com/go-qbit/model"
	"github.com/tmc/dot"
)

func (s *MySQL) GetGraphSVG() string {
	g := dot.NewGraph("Database")

	uniqRelations := map[string]map[string]*qmodel.Relation{}

	nodes := make(map[string]*dot.Node)
	for _, model := range s.models {
		n := dot.NewNode(model.GetId())
		_ = n.Set("shape", "plaintext")

		labelBuf := &bytes.Buffer{}
		labelBuf.WriteString(`<<TABLE BORDER="1" CELLBORDER="0"><TR><TD ALIGN="CENTER"><B>`)
		labelBuf.WriteString(model.GetId())
		labelBuf.WriteString(`</B></TD></TR><HR/>`)

		pkFields := map[string]bool{}
		for _, name := range model.GetPKFieldsNames() {
			pkFields[name] = true
		}

		for _, name := range model.GetFieldsNames() {
			labelBuf.WriteString(`<TR><TD ALIGN="LEFT" PORT="`)
			labelBuf.WriteString(name)
			labelBuf.WriteString(`">`)
			field := model.GetFieldDefinition(name)
			if field.IsDerivable() {
				labelBuf.WriteString(`<I>- `)
			} else if field.IsRequired() || pkFields[name] {
				labelBuf.WriteString(`∙ `)
			} else {
				labelBuf.WriteString(`∘ `)
			}
			labelBuf.WriteString(`<B>`)
			labelBuf.WriteString(field.GetId())
			labelBuf.WriteString(`</B>`)

			if !field.IsDerivable() {
				labelBuf.WriteString(`:&nbsp;`)
				labelBuf.WriteString(field.GetStorageType())
			}

			if pkFields[name] {
				labelBuf.WriteString(`<B>(PK)</B>`)
			}

			if field.IsDerivable() {
				labelBuf.WriteString(`</I>`)
			}
			labelBuf.WriteString(`</TD></TR>`)
		}

		labelBuf.WriteString(`</TABLE>>`)
		_ = n.Set("label", labelBuf.String())

		g.AddNode(n)
		nodes[model.GetId()] = n

		uniqRelations[model.GetId()] = map[string]*qmodel.Relation{}
		for _, name := range model.GetRelations() {
			if relation := model.GetRelation(name); !relation.IsBack && relation.RelationType != qmodel.RELATION_MANY_TO_MANY {
				uniqRelations[model.GetId()][relation.ExtModel.GetId()] = relation
			}
		}
	}

	for modelName, modelRelations := range uniqRelations {
		model := s.models[modelName]
		for _, relation := range modelRelations {
			if !relation.IsBack && relation.RelationType != qmodel.RELATION_MANY_TO_MANY {
				e := dot.NewEdge(nodes[model.GetId()], nodes[relation.ExtModel.GetId()])
				_ = e.Set("tailport", relation.LocalFieldsNames[0])
				_ = e.Set("headport", relation.FkFieldsNames[0])
				_ = e.Set("labelfontsize", "8")
				_ = e.Set("labeldistance", "3")
				_ = e.Set("dir", "both")

				if relation.IsRequired {
					_ = e.Set("arrowtail", "dot")
				} else {
					_ = e.Set("arrowtail", "odot")
				}

				g.AddEdge(e)
			}
		}
	}

	outBuf := &bytes.Buffer{}

	cmd := exec.Command("/usr/bin/dot", "-Tsvg")
	cmd.Stdin = bytes.NewBuffer([]byte(g.String()))
	cmd.Stdout = outBuf
	cmd.Stderr = os.Stderr
	if err := cmd.Run(); err != nil {
		panic(err)
	}

	return outBuf.String()
}
