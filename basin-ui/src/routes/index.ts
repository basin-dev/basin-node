import { BasinSDK } from "../basin-sdk/index";

export async function GET({}) {
    let basin = new BasinSDK("http://localhost:8555");
    let url = "basin://tydunn.com.twitter.followers";
    let data = await basin.read(url);
    return { status: 200, body: {data} };
}