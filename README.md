# Smasche

Each smasche worker uses a gdsm worker that connects to the gdsm manager.

```bash
docker-compose up
```

```bash
docker-compose scale workers=12
./scripts/bench.sh
```
