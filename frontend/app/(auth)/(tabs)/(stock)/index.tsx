import { StockRow } from "@/components/StockRow";
import { Candle } from "@/types/candles";
import { router, useFocusEffect } from "expo-router";
import { useCallback, useState } from "react";
import { Alert, FlatList, StyleSheet, View } from "react-native";
import { CandleService } from "@/services/candles";

export default function IndexScreen() {
    const [stocks, setStocks] = useState<Record<string, Candle[]>>({});
    const [refreshing, setRefreshing] = useState(false);

    function onGoToStock(symbol: string) {
        router.push(`/stock/${symbol}`);
    }

    async function fetchStocks() {
        try {
            setRefreshing(true);
            const response = await CandleService.stocksHistory();
            setStocks(
                Array.isArray(response.data)
                    ? { default: response.data }
                    : response.data
            );
        } catch (error) {
            if (error instanceof Error) Alert.alert("Error", error.message);
        } finally {
            setRefreshing(false);
        }
    }

    useFocusEffect(useCallback(() => { fetchStocks(); }, []));

    function renderItem({ item: symbol }: { item: string; }) {
        return (
            <StockRow
                onPress={() => onGoToStock(symbol)}
                symbol={symbol}
                candles={stocks[symbol]}
            />
        );
    }

    return (
        <FlatList
            data={Object.keys(stocks)}
            refreshing={refreshing}
            onRefresh={fetchStocks}
            keyExtractor={(symbol) => symbol}
            ItemSeparatorComponent={ItemSeparatorComponent}
            style={styles.flatList}
            renderItem={renderItem}
        />
    );
}

const ItemSeparatorComponent = () => <View style={styles.ItemSeparatorComponent} />;

const styles = StyleSheet.create({
    flatList: {
        marginBottom: 30
    },
    ItemSeparatorComponent: {
        height: 1,
        backgroundColor: "lightgray"
    }
});