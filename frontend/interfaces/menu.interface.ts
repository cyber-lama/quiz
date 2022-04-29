export interface MenuItem {
    id: {
        mainCategory: string;
    }
    pages: PageItem[];
}

export interface PageItem {
    alias:    string;
    title:    string;
    id:       number;
    category: string;
}

export enum TopLevelCategory {
    Test,
    Test1,
    Test2,
    Test3
}
