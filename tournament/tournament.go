// https://pkg.go.dev/encoding/csv#example-Reader.ReadAll
package tournament

import (
	"bufio"
	"encoding/csv"
	"fmt"
	"io"
	"sort"
	"strings"
)

var TableFormat string = "%-31v| %2v | %2v | %2v | %2v | %2v\n"

type TeamStat struct {
	name           string
	MP, W, D, L, P int
}

type TeamPresenter struct {
	name, MP, W, D, L, P string
}

func (stat *TeamStat) Win() {
	stat.MP++
	stat.W++
	stat.P += 3
}

func (stat *TeamStat) Lose() {
	stat.MP++
	stat.L++
}

func (stat *TeamStat) Draw() {
	stat.MP++
	stat.D++
	stat.P++
}

func (stat *TeamStat) ToPresenter() TeamPresenter {
	return TeamPresenter{
		stat.name,
		fmt.Sprint(stat.MP),
		fmt.Sprint(stat.W),
		fmt.Sprint(stat.D),
		fmt.Sprint(stat.L),
		fmt.Sprint(stat.P),
	}
}

func Tally(reader io.Reader, writer io.Writer) error {
	r := csv.NewReader(sanitizeInput(reader))
	r.Comma = ';'

	records, err := r.ReadAll()
	if err != nil {
		return err
	}

	for _, game := range records {
		if len(game) < 3 {
			return fmt.Errorf("Wrong format")
		}
	}

	teams, err := AggregateScores(records)
	if err != nil {
		return err
	}

	writeTeamStats(writer, teams)

	return nil
}

func sanitizeInput(reader io.Reader) io.Reader {
	scanner := bufio.NewScanner(reader)
	scanner.Split(bufio.ScanLines)
	var result []string

	for scanner.Scan() {
		if scanner.Text() == "" || strings.HasPrefix(scanner.Text(), "#") {
			continue
		}

		result = append(result, scanner.Text())
	}
	return strings.NewReader(strings.Join(result, "\n"))
}

func AggregateScores(records [][]string) ([]*TeamStat, error) {
	teams := make(map[string]*TeamStat)
	result := make([]*TeamStat, 0, len(teams))

	for _, game := range records {
		a, b, score := game[0], game[1], game[2]

		if _, ok := teams[a]; !ok {
			teams[a] = &TeamStat{name: a}
		}

		if _, ok := teams[b]; !ok {
			teams[b] = &TeamStat{name: b}
		}

		switch score {
		case "win":
			teams[a].Win()
			teams[b].Lose()
		case "draw":
			teams[a].Draw()
			teams[b].Draw()
		case "loss":
			teams[a].Lose()
			teams[b].Win()
		default:
			return nil, fmt.Errorf("unknown score")
		}
	}

	for _, value := range teams {
		result = append(result, value)
	}

	sort.Slice(result, func(i, j int) bool {
		bool := result[i].P > result[j].P

		if result[i].P == result[j].P {
			bool = result[i].name < result[j].name
		}

		return bool
	})

	return result, nil
}

func writeTeamStats(writer io.Writer, teams []*TeamStat) {
	writer.Write(formatStat(TeamPresenter{"Team", "MP", "W", "D", "L", "P"}))
	for _, team := range teams {
		writer.Write(formatStat(team.ToPresenter()))
	}
}

func formatStat(line TeamPresenter) []byte {
	result := fmt.Sprintf(TableFormat, line.name, line.MP, line.W, line.D, line.L, line.P)
	return []byte(result)
}
