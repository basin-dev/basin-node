import Basin from "@util/basin";
import {useState} from "react";


// Instead of just a hook, you might consider making a resource providing context, or a resource redux store.

export function useResource<T>(url: string) {
    let [ready, setReady] = useState(false);
    let [resource, setResource] = useState<T | undefined>(undefined);
    let [error, setError] = useState<string | undefined>(undefined);

    // First ask for permissions (what does it mean for the browser to have permissions?)

    // Then, offer the option to fill in the data if it doesn't exist, or to edit if it does

    // Then read the resource from Basin

    // Can we make this Basin.read<T>(url) and then validate to that type?
    Basin.read(url).then((resp: any) => {
        setResource(resp);
        setReady(true);
    }).catch((err: any) => {
        setError(err);
    });

    return {ready, resource, error};
}