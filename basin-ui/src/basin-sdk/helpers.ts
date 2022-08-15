export function getUserDataUrl(did: string, dataName: string): string {
    return `basin://${did}.basin.${dataName}`;
}

export function getMetadataUrl(dataUrl: string, prefix: string): string {
    let parsed = parseUrl(dataUrl);
    parsed.domain = "meta." + prefix + "." +  parsed.domain;
    return printUrl(parsed);
}

interface BasinURL {
    scheme: string;
    user: string;
    domain: string;
}

export function parseUrl(url: string): BasinURL  {
	let segs = url.split("://");

	let scheme = segs[0];

	segs = segs[1].split(".");
    let user = segs[0];
    let domain = segs.slice(1, segs.length).join(".");

	return { scheme, user, domain};
}

export function printUrl(url: BasinURL): string {
    return `${url.scheme}://${url.user}.${url.domain}`;
}