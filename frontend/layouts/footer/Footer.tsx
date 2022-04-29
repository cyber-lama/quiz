import React from "react";
import {FooterProps} from "./Footer.props";

export const Footer = ({...props}: FooterProps) => {
    return (
       <footer {...props}>Footer</footer>
    );
};
