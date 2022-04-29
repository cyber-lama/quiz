import {NextApiRequest, NextApiResponse} from "next";

export default function test(req: NextApiRequest, res:NextApiResponse) {
  if(req.method === "GET") {
    res.status(200).json({ data: "test" });
  }else{
    res.status(400).json({ error: 'Данный метод запроса не поддерживается' });
  }
}
