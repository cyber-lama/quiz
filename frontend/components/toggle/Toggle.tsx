import { motion } from "framer-motion";
import React, {memo, useContext, useEffect, useState} from "react";
import {ToggleProps} from "./Toggle.props";
import styles from "./Toggle.module.scss";
import cn from "classnames";
import {ThemeContext} from "../../contexts/theme.context";
import {isLightTheme} from "../../helpers/theme.helper";

export const Toggle = memo(({active = false, onClick, ...props}: ToggleProps) => {

    const {theme} = useContext(ThemeContext)

    const [isOn, setIsOn] = useState(active);
    useEffect(() => setIsOn(active), [active]);

    const spring = {
        type: "spring",
        stiffness: 700,
        damping: 30,
    };

    const toggleSwitch = () => {
        if(onClick) onClick();
        setIsOn(!isOn);
    };

    return (
        <div className={cn(styles.switch, {
            [styles.switch__on]: isOn,
            [styles.switch__dark]: !isLightTheme(theme),
        })} onClick={toggleSwitch} {...props}>
            <motion.div className={styles.handle} layout transition={spring} />
        </div>
    );
});
