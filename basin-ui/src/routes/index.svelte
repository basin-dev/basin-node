<script lang="ts">
    import { getUserDataUrl } from "../basin-sdk/helpers";
    import {BasinSDK} from "../basin-sdk";
    let sources: string[] = [];
    (async () => {
        let basin = new BasinSDK("http://localhost:8555");

        let did = "did:key:z6MkoYyGsB9WLBmf12RrcBdai1UPcDcyvNWcMQdRpXzzfo4H";
        let sourcesUrl = getUserDataUrl(did, "producer.sources");

        sources = await basin.read("basin://");
    })
</script>

<h1> Basin </h1>

<h3>Producer Sources</h3>
<p>{sources.join(", ")}</p>