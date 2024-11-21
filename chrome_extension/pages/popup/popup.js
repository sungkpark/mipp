import { API_V1, BASE_URL, POST, POST_IDEA } from "../../constants/api.js";

const postIdea = BASE_URL + API_V1 + POST_IDEA;

const ideaForm = document.getElementById('idea-form');

async function getCurrentCapturedUrl() {
    let queryOptions = { active: true, lastFocusedWindow: true };
    // `tab` will either be a `tabs.Tab` instance or `undefined`.
    let [tabs] = await chrome.tabs.query(queryOptions); 
    return tabs.url;
}

async function getCurrentTabDomainName() {
    let url = await getCurrentCapturedUrl()
    let domainName = url.split('://')[1].split("/")[0];
    return domainName;
}

ideaForm.addEventListener('submit', async (event) => {
    const formData = new FormData(ideaForm);

    let capturedUrl = await getCurrentCapturedUrl();
    let domainName = await getCurrentTabDomainName();

    formData.append('userName', 'poc user');
    formData.append('capturedUrl', capturedUrl);
    formData.append('domainName', domainName);
    const formDataJSON = JSON.stringify(Object.fromEntries(formData.entries()));
    const requestOptions = {
        method: POST,
        headers: {
            'Accept': 'application/json',
            'Content-Type': 'application/json',
        },
        body: formDataJSON,
    };

    fetch(postIdea, requestOptions)
        .then(response => {
            if (!response.ok) throw new Error('Network response was not ok.');
            return response.text();
        })
        .then(data => {
            console.log('Idea has been saved: ', data);
        })
        .catch(error => {
            console.error('Error: ', error);
        })    
});

document.addEventListener("DOMContentLoaded", async function () {
    const domainName = await getCurrentTabDomainName();
    document.getElementById("domain-name").textContent = domainName;
})