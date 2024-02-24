
<script setup>
import { reactive, onMounted } from 'vue'
import { CheckForUpdates } from '../../wailsjs/go/main/App'


const data = reactive({
    url: "",
    checkMessage: "Checking for updates...",
    version: "0.0",
    color: "text-cyan-200",
    disableButton: ""
})

async function checkForUpdates() {
    try {
        const result = await CheckForUpdates(); // Assuming CheckForUpdates is a function that returns a Promise
        const parsed = JSON.parse(result);
        data.disableButton = parsed.disableButton;
        data.color = parsed.color;
        data.checkMessage = parsed.message;
        data.version = parsed.version;
        data.url = parsed.url;
    } catch (error) {
        console.error('Error checking for updates:', error);
        // Handle the error appropriately
    }
}


</script>

<template>
    <div class="navbar bg-base-100">
        <div class="flex-1">
            <a class="btn btn-ghost text-xl">ICHIGO 🍓 </a>
            <div class="flex-none gap-2">
                <a class="btn btn-ghost"><svg xmlns="http://www.w3.org/2000/svg" width="16" height="16" fill="currentColor"
                        class="bi bi-house-fill" viewBox="0 0 16 16">
                        <path
                            d="M8.707 1.5a1 1 0 0 0-1.414 0L.646 8.146a.5.5 0 0 0 .708.708L8 2.207l6.646 6.647a.5.5 0 0 0 .708-.708L13 5.793V2.5a.5.5 0 0 0-.5-.5h-1a.5.5 0 0 0-.5.5v1.293z" />
                        <path d="m8 3.293 6 6V13.5a1.5 1.5 0 0 1-1.5 1.5h-9A1.5 1.5 0 0 1 2 13.5V9.293z" />
                    </svg></a>
                <a class="btn btn-ghost">Catalog</a>
                <a class="btn btn-ghost">Installed</a>
                <a class="btn btn-ghost">Uninstall</a>
            </div>
        </div>
        <div class="flex-none gap-2">
            <div class="form-control">
                <a onclick="versionModal.showModal()" @click="checkForUpdates" class="btn btn-ghost"><svg
                        xmlns="http://www.w3.org/2000/svg" width="16" height="16" fill="currentColor"
                        class="bi bi-arrow-clockwise" viewBox="0 0 16 16">
                        <path fill-rule="evenodd" d="M8 3a5 5 0 1 0 4.546 2.914.5.5 0 0 1 .908-.417A6 6 0 1 1 8 2z" />
                        <path
                            d="M8 4.466V.534a.25.25 0 0 1 .41-.192l2.36 1.966c.12.1.12.284 0 .384L8.41 4.658A.25.25 0 0 1 8 4.466" />
                    </svg>
                    Check for updates</a>
            </div>
            <div class="dropdown dropdown-end">
                <div tabindex="0" role="button" class="btn btn-ghost btn-circle avatar">
                    <div class="w-10 rounded-full">
                        <img alt="settings clog" src="../assets/images/settings.png" />
                    </div>
                </div>
            </div>
        </div>
    </div>

    <dialog id="versionModal" class="modal">
        <div class="modal-box">
            <h3 class="font-bold text-lg">Ichigo Update</h3>
            <p :class="data.color" class="py-4">{{ data.checkMessage }}</p>
            <a :href="[data.url]" target="_blank" class="py-4 text-cyan-200 underline">Release: {{ data.version }}</a>
            <div class="modal-action">
                <form method="dialog">
                    <button class="btn mr-2">Close</button>
                </form>
                <button :class="['btn btn-primary', data.disableButton]">Update</button>
            </div>
        </div>
    </dialog>
</template>