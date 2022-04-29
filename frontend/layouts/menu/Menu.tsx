import React from "react";
import {MenuProps} from "./Menu.props";

export const Menu = ({...props}: MenuProps) => {
    return (
       <section {...props}>Menu</section>
    );
};
