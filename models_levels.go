package mysql

type modelsLevels []modelLevel

type modelLevel struct {
	name  string
	level int
}

func (a modelsLevels) Len() int { return len(a) }
func (a modelsLevels) Less(i, j int) bool {
	if a[i].level == a[j].level {
		return a[i].name < a[j].name
	} else {
		return a[i].level < a[j].level
	}
}
func (a modelsLevels) Swap(i, j int) { a[i], a[j] = a[j], a[i] }

func (s *MySQL) getModelsLevels() modelsLevels {
	s.modelsMtx.RLock()
	defer s.modelsMtx.RUnlock()

	res := make(modelsLevels, 0, len(s.models))
	for name, _ := range s.models {
		res = append(res, modelLevel{
			name:  name,
			level: s.getModelLevel(name, 0, nil),
		})
	}

	return res
}

func (s *MySQL) getModelLevel(tableName string, curLevel int, visited map[string]bool) int {
	if visited == nil {
		visited = make(map[string]bool)
	}
	if visited[tableName] {
		return curLevel
	}
	visited[tableName] = true

	maxLevel := curLevel
	for _, extModel := range s.models[tableName].GetRelations() {
		relation := s.models[tableName].GetRelation(extModel)
		if relation.IsBack {
			continue
		}
		extLevel := s.getModelLevel(relation.ExtModel.GetId(), curLevel+1, visited)
		if extLevel > maxLevel {
			maxLevel = extLevel
		}
	}

	delete(visited, tableName)
	return maxLevel
}
