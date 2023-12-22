from typing import Any, Callable
import logging


def log_output(func: Callable[..., Any]) -> Callable[..., Any]:
    def wrapper(*args: Any, **kwargs: Any) -> Any:
        value = func(*args, **kwargs)
        logging.debug(f"Finished {func.__name__}({args, kwargs}) –> {value}")
        return value

    return wrapper
