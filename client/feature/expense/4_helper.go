package expense

import "fmt"


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