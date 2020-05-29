<script>
    export let click = () => {}
    export let appearance = ""
    export let type = "button"
</script>

<style>
    button {
        background: #7CD5B4;
        color: white;
        border: 0;
        border-radius: 1rem;
        font-size: 0.95em;
        font-weight: 700;
        padding: 0.8em 1.7em;
        transition: transform 200ms;
        box-shadow: 0em 0.2em 0.4em #22222212;
    }

    button:hover {
        cursor: pointer;
        transform: scale(1.05);
    }

    button:focus {
        background: #376d59;
        box-shadow: 0em 0em 2em #7CD5B4;
    }

    button.light {
        background: #9b9b9b;
    }

    button.light:focus {
        background: #424242;
        box-shadow: 0em 0em 2em rgb(199, 199, 199);
    }
</style>

<button on:click={click} class={appearance} {type}><slot /></button>
