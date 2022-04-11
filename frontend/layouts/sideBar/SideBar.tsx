import React from "react";
import {SideBarProps} from "./SideBar.props";

export const SideBar = ({...props}: SideBarProps) => {
    return (
       <div {...props}>SideBar</div>
    );
};
