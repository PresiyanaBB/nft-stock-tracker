import { IconButton } from "@/components/IconButton";
import { StockImage } from "@/components/StockImage";
import { useCandles } from "@/hooks/useCandles";
import { baseUrl } from "@/network";
import { Candle, WsCandleUpdate } from "@/types/candles";
import { CandleService } from "@/services/candles";
import { useFocusEffect, useLocalSearchParams, useNavigation } from "expo-router";
import { useCallback, useEffect, useState } from "react";
import {
    ActivityIndicator,
    Alert,
    Dimensions,
    StyleSheet,
    Text,
    View,
    ScrollView
} from "react-native";
import { LineChart as ChartKitLineChart } from "react-native-chart-kit";

export default function StockScreen() {
    const navigation = useNavigation();
    const { symbol } = useLocalSearchParams<{ symbol: string }>();

    const [visibleChart, setVisibleChart] = useState<"line" | "candlesticks">("candlesticks");
    const [isLoading, setIsLoading] = useState(true);
    const [candles, setCandles] = useState<Candle[]>([]);

    const chartWidth = Dimensions.get("screen").width - 20;
    const chartHeight = Dimensions.get("screen").height / 2;

    const {
        chartData,
        newest,
        trendingColor,
        trendingSign,
        startToEndDifference
    } = useCandles({ candles, visibleChart });

    const fetchHistory = useCallback(async () => {
        try {
            const response = await CandleService.stockCandles(symbol);
            setCandles(response.data);
        } catch (error) {
            if (error instanceof Error) Alert.alert("Error", error.message);
        } finally {
            setIsLoading(false);
        }
    }, [symbol]);

    useEffect(() => {
        navigation.setOptions({ headerTitle: symbol });
        fetchHistory();
    }, [fetchHistory, symbol]);

    useFocusEffect(
        useCallback(() => {
            const ws = new WebSocket(`${baseUrl("ws")}/ws`);
            ws.onopen = () => ws.send(symbol);
            ws.onmessage = ({ data }) => {
                const { updateType, candle } = JSON.parse(data) as WsCandleUpdate;

                if (updateType === "closed") {
                    setCandles((candles) => [...candles, candle]);
                } else {
                    setCandles((candles) => [...candles.slice(0, -1), candle]);
                }
            };

            return () => {
                ws.close();
            };
        }, [symbol])
    );

    if (isLoading) {
        return (
            <View style={styles.loader}>
                <ActivityIndicator animating size="large" />
            </View>
        );
    }

    const lineData = chartData.map(d => d.value);

    return (
        <ScrollView style={styles.container}>
            <View style={styles.innerContainer}>
                <View style={styles.imgContainer}>
                    <StockImage style={styles.img} symbol={symbol} />
                    <Text style={styles.symbol}>{symbol}</Text>
                </View>

                <View style={styles.priceContainer}>
                    <Text style={styles.price}>{"$ " + newest.close.toFixed(2)}</Text>
                    <Text style={[styles.priceStatus, { color: trendingColor }]}>
                        {trendingSign}
                        {startToEndDifference.amount.toFixed(2)}{" "}
                        ({trendingSign}{startToEndDifference.percentage.toFixed(2)}%)
                    </Text>
                </View>
            </View>

            <View style={styles.buttonsContainer}>
                <IconButton
                    name="analytics"
                    touchableOpacityStyles={{
                        backgroundColor: visibleChart === "line" ? "black" : "gray",
                    }}
                    onPress={() => setVisibleChart("line")}
                />
                <IconButton
                    name="stats-chart"
                    touchableOpacityStyles={{
                        backgroundColor: visibleChart === "candlesticks" ? "black" : "gray",
                    }}
                    onPress={() => setVisibleChart("candlesticks")}
                />
            </View>

            {visibleChart === "line" ? (
                <ChartKitLineChart
                    data={{
                        labels: [],
                        datasets: [{ data: lineData }]
                    }}
                    width={chartWidth}
                    height={chartHeight}
                    withDots={false}
                    withInnerLines={false}
                    withOuterLines={false}
                    withHorizontalLabels={false}
                    withVerticalLabels={false}
                    chartConfig={{
                        backgroundColor: "#00000000",
                        backgroundGradientFrom: "#00000000",
                        backgroundGradientTo: "#00000000",
                        color: () => trendingColor,
                    }}
                    bezier
                    style={{ marginLeft: -10 }}
                />
            ) : (
                <View style={{ height: chartHeight, justifyContent: "center", alignItems: "center" }}>
                    <Text style={{ color: "gray" }}>
                        Candlestick charts are not supported in `react-native-chart-kit`.
                    </Text>
                </View>
            )}
        </ScrollView>
    );
}

const styles = StyleSheet.create({
    loader: {
        flex: 1,
        justifyContent: "center",
        alignItems: "center",
    },
    container: {
        flex: 1,
    },
    innerContainer: {
        flexDirection: "row",
        justifyContent: "space-between",
        alignItems: "center",
        marginBottom: 20,
        padding: 20,
    },
    imgContainer: {
        flexDirection: "row",
        justifyContent: "center",
        alignItems: "center",
        gap: 10,
    },
    img: {
        width: 70,
        height: 70,
    },
    symbol: {
        fontSize: 25,
        fontWeight: "bold",
    },
    priceContainer: {
        justifyContent: "center",
        alignItems: "flex-end",
        alignSelf: "center",
        gap: 5,
    },
    buttonsContainer: {
        flexDirection: "row",
        justifyContent: "center",
        gap: 30,
        marginBottom: 20,
    },
    price: {
        fontSize: 26,
        fontWeight: "bold",
    },
    priceStatus: {
        fontSize: 15,
        fontWeight: "semibold",
    },
});
