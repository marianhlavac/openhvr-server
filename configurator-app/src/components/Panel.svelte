<script>
    export var status = 'none'
</script>

<style>
    .panel {
        position: relative;
        background: white;
        padding: 2rem 1.5rem;
        border-radius: 1em;
        margin-bottom: 2em;
        box-shadow: 0em 0.2em 0.4em #22222208;
    }

    @keyframes loading-stripe {
        from { background-position-x: 200%; }
        to { background-position-x: 0%; }
    }

    .panel.loading {
        background: rgb(255,255,255);
        background: linear-gradient(93deg, rgba(255,255,255,0) 0%, rgba(255,255,255,1) 76%, rgba(255,255,255,0) 100%);
        background-size: 200%;
        animation: 1s linear infinite loading-stripe;
    }
</style>

<div class="panel {status}">
    <slot />
</div>
