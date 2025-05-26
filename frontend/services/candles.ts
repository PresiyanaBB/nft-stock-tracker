import { CandleListResponse } from "@/types/candles";
import { Api } from "./api";

async function stocksHistory(): Promise<CandleListResponse> {
    return Api.get("/candle/stocks-history");
}

async function stockCandles(symbol: string): Promise<CandleListResponse> {
    return Api.get(`/candle/stock-candles/${symbol}`);
}


const CandleService = {
    stocksHistory,
    stockCandles,
};

export { CandleService };