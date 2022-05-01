import {DetailedHTMLProps, HTMLAttributes, ReactNode} from "react";

export interface ThemeLayoutProps extends DetailedHTMLProps<HTMLAttributes<HTMLDivElement>, HTMLDivElement>{
    children: ReactNode
}