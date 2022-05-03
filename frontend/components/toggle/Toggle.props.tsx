import {DetailedHTMLProps, HTMLAttributes} from "react";

export interface ToggleProps extends DetailedHTMLProps<HTMLAttributes<HTMLDivElement>, HTMLDivElement>{
    active?: boolean,
    onClick?: () => void
}