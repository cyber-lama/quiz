export const setLocalStorage = (name: string, val: string | Record<string, unknown>) => {
    if (typeof window !== 'undefined') {
        localStorage?.setItem(name, JSON.stringify(val));
    }
};

export const getLocalStorage = (name:string) => {
    let out = null;
    if (typeof window === 'undefined') return out;

    const saved = localStorage.getItem(name);
    if(saved){
        out = JSON.parse(saved);
    }
    return out;
}