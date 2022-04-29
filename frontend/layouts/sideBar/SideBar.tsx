import React from "react";
import {SideBarProps} from "./SideBar.props";

export const SideBar = ({...props}: SideBarProps) => {
    return (
       <section {...props}>SideBar</section>
    );
};
