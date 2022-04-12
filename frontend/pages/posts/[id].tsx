import React, {FunctionComponent} from "react";
import {GetStaticPaths, GetStaticProps, GetStaticPropsContext} from "next";
import axios from "axios";
import {WithMainLayout} from "../../layouts";
import {PostItem} from "../../interfaces/post.interface";
import {ParsedUrlQuery} from "querystring";

const Post = ({post}:PostProps) => {
    if(!post) return <p>Ожидайте загрузки</p>;
    return (
       <>
           <p>{post.title}</p>
           <p>{post.body}</p>
       </>
    );
};

export default WithMainLayout(Post);
// getStaticPaths Нужен, чтобы получить пути, по которым должен пройти генератор статики на этапе сборки.
export const getStaticPaths: GetStaticPaths = async () => {
    const { data: posts } = await axios.get<PostItem[]>("https://jsonplaceholder.typicode.com/posts");
    const paths = posts.map(post => ({
        params: {id: post.id.toString()}
    }));
    return{paths, fallback: true};
};

export const getStaticProps: GetStaticProps<PostProps> = async ({params}: GetStaticPropsContext<ParsedUrlQuery>) => {
    if(!params){
        return {
            notFound: true
        };
    }
    let  post;
    try {
        const { data } = await axios.get<PostItem>(`https://jsonplaceholder.typicode.com/posts/${params.id}`);
        post = data;
    }catch (e) {
        return {
            notFound: true
        };
    }
    return {
        props: {
            post
        }
    };
};

interface PostProps extends Record<string, unknown> {
    post: PostItem;
}
