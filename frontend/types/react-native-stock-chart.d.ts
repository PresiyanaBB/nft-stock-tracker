declare module "react-native-stock-chart" {
    import React from "react";
    import { ViewStyle } from "react-native";

    export interface CandlestickChartProps {
        data: {
            timestamp: number;
            open: number;
            high: number;
            low: number;
            close: number;
        }[];
        width: number;
        height: number;
        candleWidth?: number;
        candleColorPositive?: string;
        candleColorNegative?: string;
        style?: ViewStyle;
    }

    export default class CandlestickChart extends React.Component<CandlestickChartProps> { }
}
