export const getDomains = async () => {
    const res = await fetch(`${process.env.MIPP_BACKEND_URL}/domains`, {
        method: 'GET',
        headers: { 'Content-Type': 'application/json' },
    })
    return await res.json()
}