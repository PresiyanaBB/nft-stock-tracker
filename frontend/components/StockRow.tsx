import { Candle } from "@/types/candles";
import { StyleSheet, Text, TouchableOpacity, View } from "react-native";
import { StockImage } from "./StockImage";
import { LineChart as ChartKitLineChart } from "react-native-chart-kit";
import { useCandles } from "@/hooks/useCandles";

interface Props {
    symbol: string;
    candles: Candle[];
    onPress: () => void;
}

export function StockRow({ candles, symbol, onPress }: Props) {
    const {
        chartData,
        newest,
        trendingSign,
        trendingColor,
        startToEndDifference
    } = useCandles({ candles, visibleChart: 'line' });

    const lineData = chartData.map(item => ({
        open: item.open,
        high: item.high,
        low: item.low,
        close: item.close,
        timestamp: item.timestamp,
    }));

    return (
        <TouchableOpacity style={styles.container} onPress={onPress}>
            <View style={styles.imageContainer}>
                <StockImage style={styles.img} symbol={symbol} />
                <Text style={styles.symbol}>{symbol}</Text>
            </View>

            <ChartKitLineChart
                data={{
                    labels: [], // Optional, or use timestamps
                    datasets: [{ data: lineData.map(item => item.close) }],
                }}
                width={100}
                height={100}
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
                    propsForBackgroundLines: {
                        strokeDasharray: "", // solid lines
                    },
                }}
                bezier
                style={{ paddingRight: 0 }}
            />

            <View style={styles.priceContainer}>
                <Text style={styles.price}>
                    {"$ " + newest.close.toFixed(2)}
                </Text>
                <Text style={[styles.priceStatus, { color: trendingColor }]}>
                    {trendingSign}
                    {startToEndDifference.amount.toFixed(2)}
                    {" "}
                    ({trendingSign}{startToEndDifference.percentage.toFixed(2) + "%"})
                </Text>
            </View>
        </TouchableOpacity>
    );
}

const styles = StyleSheet.create({
    container: {
        flexDirection: 'row',
        justifyContent: "space-between",
        alignItems: "center",
        paddingHorizontal: 10,
    },
    imageContainer: {
        flexDirection: "row",
        justifyContent: "center",
        alignItems: "center",
        gap: 10
    },
    img: {
        width: 60,
        height: 60,
        margin: 10,
        borderRadius: 30,
    },
    symbol: {
        fontSize: 18,
        fontWeight: 'bold',
    },
    priceContainer: {
        justifyContent: "center",
        alignItems: "flex-end",
        gap: 5,
    },
    price: {
        fontSize: 22,
        fontWeight: "bold"
    },
    priceStatus: {
        fontSize: 15,
        fontWeight: "semibold"
    }
});
