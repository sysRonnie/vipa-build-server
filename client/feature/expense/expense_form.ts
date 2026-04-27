export function validateExpenseForm(form: HTMLFormElement): boolean {
    const data = new FormData(form);

    let hasAtLeastOneName = false;

    for (let i = 0; i < 5; i++) {
        const name = data.get(`clients[${i}].name`)?.toString().trim();
        const email = data.get(`clients[${i}].email`)?.toString().trim();
        const phone = data.get(`clients[${i}].phone`)?.toString().trim();
        const address = data.get(`clients[${i}].address`)?.toString().trim();

        const hasAnyField = name || email || phone || address;

        // 🚨 Rule 1: partial row without name → invalid
        if (hasAnyField && !name) {
            alert(`Row ${i + 1}: Name is required if other fields are filled`);
            return false;
        }

        // ✅ Track if at least one valid row exists
        if (name) {
            hasAtLeastOneName = true;
        }
    }

    // 🚨 Rule 2: must have at least one name
    if (!hasAtLeastOneName) {
        alert("At least one client (with name) is required");
        return false;
    }

    return true;
}
document.body.addEventListener("htmx:beforeRequest", (e: any) => {
    const form = e.target as HTMLFormElement;

    if (form.id !== "expense-form") return;

    const isValid = validateExpenseForm(form);

    if (!isValid) {
        e.preventDefault(); // ✅ stops HTMX request
    }
});