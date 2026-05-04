package expense

import (
	"fmt"
	"strconv"
)


func first(data map[string][]string, key string) string {
    if v, ok := data[key]; ok && len(v) > 0 {
        return v[0]
    }
    return ""
}
func parseClients(data map[string][]string) []ExpenseFormModelClient {
    clients := []ExpenseFormModelClient{}

    for i := 0; ; i++ {
        nameKey := fmt.Sprintf("clients[%d].name", i)

        names := data[nameKey]
        if len(names) == 0 || names[0] == "" {
            break
        }

        client := ExpenseFormModelClient{
            Name:    names[0],
            Email:   first(data, fmt.Sprintf("clients[%d].email", i)),
            Phone:   first(data, fmt.Sprintf("clients[%d].phone", i)),
            Address: first(data, fmt.Sprintf("clients[%d].address", i)),
        }

        clients = append(clients, client)
    }

    return clients
}




func ParseClientFormSingleRow(data map[string][]string) ExpenseFormModelClient {
	return ExpenseFormModelClient{
		Name:    first(data, "name"),
		Email:   first(data, "email"),
		Phone:   first(data, "phone"),
		Address: first(data, "address"),
	}
}

func BuildExpenseTableModel(clients []ExpenseFormModelClient) ExpenseTableModel {
	columns := []Column{
		{Key: "id", Label: "ID"},
		{Key: "name", Label: "Name"},
		{Key: "email", Label: "Email"},
		{Key: "phone", Label: "Phone"},
		{Key: "address", Label: "Address"},
	}

	var rows []map[string]string

	for _, c := range clients {
		row := map[string]string{
			"id":      strconv.Itoa(c.Id),
			"name":    c.Name,
			"email":   c.Email,
			"phone":   c.Phone,
			"address": c.Address,
		}
		rows = append(rows, row)
	}

	return ExpenseTableModel{
		Columns: columns,
		Rows:    rows,
	}
}