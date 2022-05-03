import React, {FunctionComponent} from "react";
import {MainLayoutProps} from "./MainLayout.props";
import {SideBar} from "../sideBar/SideBar";
import {Header} from "../header/Header";
import {Footer} from "../footer/Footer";
import {IAppContext} from "../../contexts/app.context";

const MainLayout = ({children}: MainLayoutProps) => {
    return (
        <>
            <Header/>
            <div>
                <SideBar/>
                <div>
                    {children}
                </div>
            </div>
            <Footer/>
        </>
    );
};

export const WithMainLayout =<T extends Record<string, unknown> & IAppContext> (Component: FunctionComponent<T>) => {

    return function WithMainLayoutComponent(props: T):JSX.Element{
        return (
                <MainLayout>
                    <Component {...props}/>
                </MainLayout>
        );
    };
};