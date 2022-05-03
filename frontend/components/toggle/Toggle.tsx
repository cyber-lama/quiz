import { motion } from "framer-motion";
import React, {useState} from "react";
import {ToggleProps} from "./Toggle.props";
import styles from "./Toggle.module.scss";
import cn from "classnames";

export const Toggle = ({active = false, onClick, ...props}: ToggleProps) => {
    const spring = {
        type: "spring",
        stiffness: 700,
        damping: 30
    };

    const [isOn, setIsOn] = useState(active);
    const toggleSwitch = () => {
        if(onClick) onClick();
        setIsOn(!isOn);
    };

    return (
        <div className={cn(styles.switch, {
            [styles.switch__on]: isOn
        })} onClick={toggleSwitch} {...props}>
            <motion.div className={styles.handle} layout transition={spring} />
        </div>
    );
};
