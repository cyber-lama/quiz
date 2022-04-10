import {RatingProps} from "./Rating.props";
import {useEffect, useState, KeyboardEvent} from "react";
import RatingIcon from "./raiting.svg";
import styles from "./Rating.module.scss";
import cn from "classnames";

export const Rating = ({isEditable = false, rating, setRating, className, ...props}:RatingProps): JSX.Element => {
    const [ratingArray, setRatingArray] = useState<JSX.Element[]>(new Array(5).fill(<></>));

    useEffect(()=>{
        constructRating(rating);
    }, [rating]);
    const changeDisplay = (i:number) => {
        if(!isEditable) return;
        constructRating(i);
    };
    const onClick = (i:number) => {
        if(!isEditable || !setRating) return;
        setRating(i);
    };
    const handleSpace = (i:number, e: KeyboardEvent<SVGElement>) => {
        console.log(e.code);
        if(e.code !== "Space" || !setRating) return;
        setRating(i);
    };
    const constructRating = (currentRating:number) => {
        const updateArray = ratingArray.map((r, i) => {
            return (
                <RatingIcon className={cn(styles.star, {
                    [styles.filled]: i < currentRating,
                    [styles.editable] : isEditable
                })}
                            onMouseEnter={() => changeDisplay(i + 1)}
                            onMouseLeave={() => changeDisplay(rating)}
                            onClick={() => onClick(i + 1)}
                            tabIndex={isEditable? 0 : -1}
                            onKeyDown={(e: KeyboardEvent<SVGElement>) => isEditable && handleSpace(i + 1, e)}
                />
            );
        });
        setRatingArray(updateArray);
    };

    return (
        <div className={className} {...props}>
            {ratingArray.map((r,i) => <span key={i}>{r}</span>)}
        </div>
    );
};