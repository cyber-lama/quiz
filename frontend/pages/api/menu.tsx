// Next.js API route support: https://nextjs.org/docs/api-routes/introduction

import {NextApiRequest, NextApiResponse} from "next";
import menuItems from "../../mocks/menu.mock";

export default function getMenuItems(req: NextApiRequest, res:NextApiResponse) {

  if(req.method === "GET") {
    res.status(200).json({ data: menuItems });
  }else{
    res.status(400).json({ error: 'Данный метод запроса не поддерживается' });
  }
}
