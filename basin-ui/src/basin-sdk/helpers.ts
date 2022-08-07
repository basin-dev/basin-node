export function getUserDataUrl(did: string, dataName: string): string {
    return `basin://${did}.basin.${dataName}`;
}