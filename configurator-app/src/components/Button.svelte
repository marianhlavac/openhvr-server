<script>
    export var click = () => {}
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
</style>

<button on:click={click}><slot /></button>
