const API_ENDPOINT = 'http://127.0.0.1:47023/v1'

export async function fetchDeviceList() {
    const res = await fetch(`${API_ENDPOINT}/devices`)
    return await res.json()
}

export async function fetchDevice(id) {
    const res = await fetch(`${API_ENDPOINT}/devices/${id}`)
    return await res.json()
}

export async function updateDevice(id, device) {
    const url = id !== null ? `${API_ENDPOINT}/devices/${id}` : `${API_ENDPOINT}/devices/`
    const res = await fetch(url, {
        method: id !== null  ? 'put' : 'post',
        body: JSON.stringify(device)
    })
    return res.ok
}

export async function deleteDevice(id) {
    const url = `${API_ENDPOINT}/devices/${id}`
    const res = await fetch(url, { method: 'delete' })
    return res.ok
}

export async function fetchDrivers() {
    const res = await fetch(`${API_ENDPOINT}/devices/drivers`)
    return await res.json()
}

export async function fetchTypes() {
    const res = await fetch(`${API_ENDPOINT}/effects/types`)
    return await res.json()
}
