async function getCurrentTabDomainName() {
    let queryOptions = { active: true, lastFocusedWindow: true };
    // `tab` will either be a `tabs.Tab` instance or `undefined`.
    let [tabs] = await chrome.tabs.query(queryOptions);
    let domainName = tabs.url.split('://')[1].split("/")[0];
    console.log(domainName);
    return domainName;
}

// chrome.tabs.query({ currentWindow: true, active: true }, function (tabs) {
//     console.log(tabs);
// })

document.addEventListener("DOMContentLoaded", async function () {
    const domainName = await getCurrentTabDomainName();
    document.getElementById("domain-name").textContent = domainName;
})