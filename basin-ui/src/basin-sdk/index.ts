/** The Basin SDK lets web developers retrieve Basin resources as easily as with fetch. It also has capabilities for type generation. */
/** Later...the SDK should be able to optionally submit appeals if an object with wrong schema (or nothing at all) is returned. */

export class BasinSDK {
    /** The HTTP(S) URL of the node you are using as a gateway to the network. */
    gatewayUrl: string;
    constructor(_gatewayUrl: string) {
        this.gatewayUrl = _gatewayUrl;
    }

    /** Read a Basin resource by URL. If you are not subscribed you can do so automatically. */
    async read(url: string): Promise<any> {
        let fullUrl = `${this.gatewayUrl}/read?url=${encodeURIComponent(url)}`;
        let resp = await fetch(fullUrl, {
            method: "GET",
            mode: "no-cors",
            headers: {
                'Access-Control-Allow-Origin':'*',
                "Access-Control-Allow-Methods": "GET, OPTIONS, POST, PUT"
            }
        });
        if (typeof window !== "undefined") return;
        let base64 = await resp.json();
        let buffer = Buffer.from(base64, "base64");
        return buffer.toString();
    };
    // write: (url: string, val: any) => Promise<boolean>;
}

/**
 * NOTES:
 * The SDK needs to be able to sign messages. It should be passes a DID, and then the browser extension or something should deal with signing.
 * Want to build in not just the HTTP interface, but the subscription request popup, type generation, and forms.
 */