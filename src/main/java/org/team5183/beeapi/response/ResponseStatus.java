package org.team5183.beeapi.response;

public enum ResponseStatus {
    SUCCESS("success"),
    ERROR("error");

    private final String status;

    ResponseStatus(String status) {
        this.status = status;
    }

    public String getStatus() {
        return status;
    }
}
