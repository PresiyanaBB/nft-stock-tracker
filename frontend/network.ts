import { Platform } from "react-native";

export const baseUrl = (scheme: "http" | "ws") => {
    const PORT = 8081
    const HOST = Platform.OS === "android" ? "192.168.43.150" : "localhost"

    return `${scheme}://${HOST}:${PORT}`
}