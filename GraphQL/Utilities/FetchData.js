import { PopUp } from "./Popups.js";

export const FetchQL = async (query, token) => {
    try {
        const response = await fetch("https://learn.zone01oujda.ma/api/graphql-engine/v1/graphql", {
            method: "POST",
            headers: {
                "Content-Type": "application/json",
                "Authorization": `Bearer ${token}`,
            },
            body: JSON.stringify({
                query: query
            })
        })
        return response.json()
    } catch (err) {
        PopUp(err)
    }
}