import {MenuItem} from "../interfaces/menu.interface";

const menuItems: MenuItem[] = [{
    id: 0,
    isOpened: false,
    pages: [{
        alias: "test-subtitle",
        title: "Тестовый подзаголовок",
        id: 0,
        category: "Test"
    },{
        alias: "test-subtitle1",
        title: "Тестовый подзаголовок 1",
        id: 1,
        category: "Test"
    },{
        alias: "test-subtitle2",
        title: "Тестовый подзаголовок2",
        id: 2,
        category: "Test"
    }]
}];

export default menuItems;
