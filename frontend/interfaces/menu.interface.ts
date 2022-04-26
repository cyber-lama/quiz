export interface PageItem {
    alias: string;
    title: string;
    id: number;
    category: string;
}

export interface MenuItem {
    id: number
    isOpened?: boolean;
    pages: PageItem[];
}


export enum TopLevelCategory {
    Test,
    Test1,
    Test2,
    Test3
}
